import Button from "../../../components/button";
import styles from "../../shared/style/modal.module.css";
import { useDeleteTeacherMutation } from "../api/useDeleteTeacherMutation";

export function DeleteModal({ deleteTeacher, isOpen, original }) {
  const deleteTeacherMutation = useDeleteTeacherMutation();
  const id = original.id;

  if (!isOpen) {
    return null;
  }
  function deleteTeacherHandle() {
    deleteTeacherMutation.mutateAsync({ id });
    deleteTeacher(true);
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <p className={styles['delete-text']}>Вы уверены, что хотите удалить {original.name}?</p>
        <div className={styles["modal-delete"]}>
          <Button
            styleFeature='close'
            type="submit"
            onClick={() => deleteTeacherHandle()}
          >
            Удалить
          </Button>
          <Button onClick={deleteTeacher}>Отмена</Button>
        </div>
      </div>
    </div>
  );
}
