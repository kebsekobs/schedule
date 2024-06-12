import Button from "../../../components/button";
import { CloseSvg } from "../../../components/close-svg";
import styles from "../../shared/style/modal.module.css";
import { useDeleteClassroomMutation } from "../api/useDeleteClassroomMutation";

export function DeleteModal({ deleteClassroom, isOpen, original }) {
  const deleteClassroomMutation = useDeleteClassroomMutation();
  const id = original.id;

  if (!isOpen) {
    return null;
  }
  function deleteGroup() {
    deleteClassroomMutation.mutateAsync({ id });
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <p className={styles['delete-text']}>Вы уверены, что хотите удалить {original.id}?</p>
        <div className={styles["modal-delete"]}>
          <Button
            styleFeature='close'
            type="submit"
            onClick={() => deleteGroup()}
          >
            Удалить
          </Button>
          <Button onClick={deleteClassroom}>Отмена</Button>
        </div>
      </div>
    </div>
  );
}
